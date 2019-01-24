package com.example.renosyahputra.grpcservicelib.ModelCasting;

import auth.Auth;
import com.example.renosyahputra.grpcservicelib.model.UserDataModel;

public class UserDataModelCasting {

    public static UserDataModel toUserData(Auth.userData userDataModel){
        return new UserDataModel(userDataModel.getId(),
                userDataModel.getName(),
                userDataModel.getEmail(),
                userDataModel.getUsername());
    }

    public static Auth.userData toUserDataGRPC(UserDataModel userDataModel){
        return Auth.userData.newBuilder()
                .setId(userDataModel.Id)
                .setName(userDataModel.Name)
                .setEmail(userDataModel.Email)
                .setPassword("")
                .setUsername(userDataModel.Username)
                .build();
    }
}
