library user_register;

import 'dart:html';

import 'package:angular/angular.dart';

import 'package:user_management/service/users.dart';
import 'package:user_management/model/user.dart';

@Controller(selector: '[user-register]', publishAs: 'form')
class RegisterFormController {

  String displayName;
  String email;
  String password;

  Map<String, String> inputErrors = new Map();

  bool isInputError(field) => inputErrors.containsKey(field);

  String errorMessage;

  bool isError() => errorMessage != '' && errorMessage != null;

  final UsersService _usersService;
  final Window _window;

  RegisterFormController(this._usersService, this._window);

  void registerSubmit() {
    errorMessage = '';
    inputErrors.clear();
    var query = _window.location.search;
    if (query.length == 0) {
      query = '?';
    }
    var params = Uri.splitQueryString(query.substring(1));
    var user = new User(displayName, email, password);
    _usersService.registerUser(user, params['continue']).then((RegisterResponse response) {
      if (!response.valid) {
        response.errors.forEach((error) {
          inputErrors[error.name] = error.errorMessage;
        });
      } else {
        _window.location.assign(response.redirectUri);
      }
    }).catchError((error) {
      errorMessage = 'There was an error. Please try again.';
    });
  }
}
