package com.example.renosyahputra.grpcservicelib.model;

import java.io.Serializable;

public class UserDataModel implements Serializable {
    public String Id = "";
    public String Name = "";
    public String Email = "";
    public String Username  = "";

    public UserDataModel() {
        super();
    }

    public UserDataModel(String id, String name, String email, String username) {
        Id = id;
        Name = name;
        Email = email;
        Username = username;
    }
}
