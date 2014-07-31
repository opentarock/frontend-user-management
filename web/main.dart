library user_managment;

import 'package:angular/angular.dart';
import 'package:angular/application_factory.dart';

import 'package:user_management/controller/user_register.dart';
import 'package:user_management/service/users.dart';

class UserRegisterModule extends Module {
  UserRegisterModule() {
    bind(UserRegisterController);
    bind(UsersService);
  }
}

void main() {
  applicationFactory()
    .addModule(new UserRegisterModule())
    .run();
}
