package cmd

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/renosyah/CustomAuth/Golang/router"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"os/signal"
	"syscall"
	"context"
	"net"
	"github.com/renosyah/CustomAuth/Golang/auth"
)

var (
	cfgFile   string
	dbPool    *sql.DB
)

var rootCmd = &cobra.Command{
	Use: "CustomAuth",
	Run: func(cmd *cobra.Command, args []string) {

		ctx, cancel := context.WithCancel(SignalContext(context.Background()))
		defer cancel()

		initDB()

		router.Init(dbPool)
		auth.Init(dbPool)

		t := &auth.Auth_Server{
			Broadcast: make(chan auth.CallbackData),
			ClientStreams: make(map[string]chan auth.CallbackData),
		}

		t_holder := &router.Auth_Holder{
			Auth:t,
		}

		lis, err := net.Listen("tcp", fmt.Sprint(":", viper.GetInt("app.port_grpc")))
		if err != nil {
			fmt.Println("failed to serve: ", err.Error())
		}

		s := grpc.NewServer()

		auth.RegisterAuthServiceServer(s,t)

		reflection.Register(s)


		r := mux.NewRouter()
		http.Handle("/", r)

		r.HandleFunc("/register",router.RegisterUserPage)
		r.HandleFunc("/",router.LoginUserPage)

		r.HandleFunc("/req/register",router.RegisterUser)
		r.HandleFunc("/req/login",t_holder.LoginUser)

		http.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.Dir("data"))))
		http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
		http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
		http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("fonts"))))

		go t_holder.Auth.MakeBroadcast(ctx)

		go func() {

			if err := http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("app.port")), nil); err != nil {
				fmt.Println("failed to serve: ", err.Error())
			}
			cancel()
		}()
		go func() {

			if err := s.Serve(lis); err != nil {
				fmt.Println("failed to serve: ", err.Error())
			}
			cancel()
		}()

		<-ctx.Done()
		close(t.Broadcast)
		s.GracefulStop()
	},
}

func Execute(){
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}
func initConfig() {
	viper.SetConfigType("toml")
	if cfgFile != "" {

		viper.SetConfigFile(cfgFile)
	} else {

		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.AddConfigPath("/etc/CustomAuth")
		viper.SetConfigName(".CustomAuth")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func initDB() {

	Host:= viper.GetString("database.host")
	Port:= viper.GetInt("database.port")
	Username:= viper.GetString("database.username")
	Password:= viper.GetString("database.password")
	DBName:= viper.GetString("database.name")

	var db, err = sql.Open("mysql", fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s`,Username,Password,Host,Port,DBName))
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)

	}
	dbPool = db
}

func SignalContext(ctx context.Context) context.Context {

	ctx, cancel := context.WithCancel(ctx)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {

		<-sigs
		signal.Stop(sigs)
		close(sigs)
		cancel()
	}()
	return ctx
}