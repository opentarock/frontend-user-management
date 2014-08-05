library user_service_spec;

import 'dart:async' as asy;
import 'dart:convert';

import 'package:guinness/guinness.dart';
import 'package:unittest/unittest.dart' as unit;
import 'package:di/di.dart';
import 'package:angular/angular.dart';
import 'package:angular/mock/module.dart';

import 'package:user_management/service/users.dart';
import 'package:user_management/model/user.dart';

void main() {
  describe('RegisterFormController', () {
    beforeEach(() {
      setUpInjector();
      module((Module m) {
        m.bind(UsersService);
      });
    });

    afterEach(tearDownInjector);

    describe('registerUser', () {
      it('should return redirect uri on success', inject(
        (UsersService s, MockHttpBackend http) {
          var user = new User('display name', 'email@example.com', 'pass');
          var response = new RegisterResponse.empty()
                            ..valid = true
                            ..redirectUri = '/redirect';
          http.expectPOST('/api/user', JSON.encode(user))
              .respond('{"valid":true, "redirect_uri": "/redirect"}');

          // TODO: make it working
//          var future = s.registerUser(user).then((r) {
//            expect(r.valid).toBeFalse();
//            expect(r.redirectUri).toBe('/redirec');
//          }, onError: (_) { throw 'failed'; });
//          unit.expect(future, unit.completes);
        }
      ));
    });
  });
}