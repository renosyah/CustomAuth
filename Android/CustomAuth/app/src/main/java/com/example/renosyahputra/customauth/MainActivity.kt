package com.example.renosyahputra.customauth

import android.os.Bundle
import android.support.v7.app.AppCompatActivity
import android.view.View
import com.example.renosyahputra.grpcservicelib.dialogAuth.DialogAuth
import com.example.renosyahputra.grpcservicelib.model.UserDataModel
import kotlinx.android.synthetic.main.activity_main.*
import java.util.*


class MainActivity : AppCompatActivity(),View.OnClickListener, DialogAuth.OnAuthListener {


    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        buttonLogin.setOnClickListener(this)

    }

    override fun onClick(v: View?) {
        when (v) {
            buttonLogin -> {

                DialogAuth.newInstanceWithContext(this@MainActivity)
                    .setOnAuthListener(this)
                    .show()

            }
        }
    }

    override fun onAuthSucces(user: UserDataModel) {
        result.text = "Name : ${user.Name}\nEmail : ${user.Email}\nUsername : ${user.Username}"
    }

    override fun onAuthFail(Errors: ArrayList<String>) {
        errorResult.text = "Error Messages : ${Errors.toString()}"
    }
}
