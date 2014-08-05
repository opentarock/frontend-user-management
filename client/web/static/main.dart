library user_managment;

import 'package:angular/angular.dart';
import 'package:angular/application_factory.dart';

import 'package:user_management/controller/register_form.dart';
import 'package:user_management/service/users.dart';

class UserRegisterModule extends Module {
  UserRegisterModule() {
    bind(RegisterFormController);
    bind(UsersService);
  }
}

void main() {
  applicationFactory()
    .addModule(new UserRegisterModule())
    .run();
}
