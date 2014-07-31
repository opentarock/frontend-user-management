library user_service;

import 'dart:async';
import 'dart:convert';

import 'package:angular/angular.dart';

import 'package:user_management/model/user.dart';

@Injectable()
class UsersService {
  final Http _http;

  UsersService(this._http);

  Future<RegisterSuccess> registerUser(User user) {
    return _http.post('/register', JSON.encode(user))
                .then((HttpResponse r) => new RegisterSuccess.fromJson(r.data))
                .catchError((HttpResponse r) {
                  if (r.status == 400) {
                    throw new RegistrationException.fromJson(r.data);
                  } else {
                    throw r;
                  }
                }, test: (e) => e is HttpResponse);
  }
}

class RegisterSuccess {
  int userId;

  RegisterSuccess(this.userId);

  RegisterSuccess.fromJson(Map<String, dynamic> json) : this(int.parse(json['user_id']));
}

class RegistrationException implements Exception {
  final String type;
  final String message;

  const RegistrationException(this.type, this.message);

  RegistrationException.fromJson(Map<String, dynamic> json) : this(json['type'], json['message']);

  String toString() => "RegistrationException: $message";
}

isRegistrationError(error) => error is RegistrationException;