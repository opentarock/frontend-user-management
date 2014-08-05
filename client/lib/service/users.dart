library user_service;

import 'dart:async';
import 'dart:convert';

import 'package:angular/angular.dart';

import 'package:user_management/model/user.dart';

@Injectable()
class UsersService {
  final Http _http;

  UsersService(this._http);

  Future<RegisterResponse> registerUser(User user, [String redirectUri]) {
    String apiUrl = '/api/user';
    if (redirectUri != null) {
      apiUrl += '?redirect_uri=' + Uri.encodeComponent(redirectUri);
    }
    return _http.post(apiUrl, JSON.encode(user))
                .then((HttpResponse r) => new RegisterResponse.fromJson(r.data),
                      onError: (HttpResponse e) {
                        if (e.status == 400) {
                          return new RegisterResponse.fromJson(JSON.decode(e.data));
                        } else {
                          throw e;
                        }
                      });
  }
}

class InputError {
  final String name;
  final String errorMessage;

  const InputError(this.name, this.errorMessage);

  InputError.fromJson(Map<String, dynamic> json) : this(json['name'], json['error_message']);
}

class RegisterResponse {
  String redirectUri;
  bool valid;
  List<InputError> errors;

  RegisterResponse.empty();

  RegisterResponse(this.redirectUri, this.valid, this.errors);

  RegisterResponse.fromJson(Map<String, dynamic> json) {
    valid = json['valid'];
    if (!valid) {
      var errorsList = json['errors'] as List<Map<String, dynamic>>;
      errors = [];
      errorsList.forEach((e) {
        errors.add(new InputError.fromJson(e));
      });
    } else {
      redirectUri = json['redirect_uri'];
    }
  }
}
