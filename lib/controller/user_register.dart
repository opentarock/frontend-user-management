library user_register;

import 'dart:html';

import 'package:angular/angular.dart';

import 'package:user_management/service/users.dart';
import 'package:user_management/model/user.dart';

@Controller(selector: '[user-register]', publishAs: 'ctrl')
class UserRegisterController {

  String displayName;
  String email;
  String password;

  String errorMessage;

  bool isError() => errorMessage != '' && errorMessage != null;

  UsersService _usersService;
  Window _window;

  UserRegisterController(this._usersService, this._window);

  void registerSubmit() {
    errorMessage = '';
    var user = new User(displayName, email, password);
    _usersService.registerUser(user).then((RegisterSuccess result) {
      _window.location.assign('/user');
    }).catchError((RegistrationException error) {
      errorMessage = error.message;
    }, test: (e) => isRegistrationError(e))
    .catchError((error) {
      errorMessage = 'There was an error. Please try again.';
    });
  }
}
