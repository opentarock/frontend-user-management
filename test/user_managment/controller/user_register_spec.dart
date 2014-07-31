library user_register_spec;

import 'dart:async';
import 'dart:html';

import 'package:guinness/guinness.dart';
import 'package:mockito/mockito.dart';
import 'package:mock/mock.dart' as m;
import 'package:di/di.dart';
import 'package:angular/angular.dart';
import 'package:angular/mock/module.dart';

import 'package:user_management/controller/user_register.dart';
import 'package:user_management/service/users.dart';
import 'package:user_management/model/user.dart';

@proxy
class MockUsersService extends Mock implements UsersService {
  noSuchMethod(Invocation invocation) => super.noSuchMethod(invocation);
}

void main() {
  describe('UserRegisterController', () {
    beforeEach(() {
      setUpInjector();
      module((Module m) {
        m.bind(UserRegisterController);
        var usersService = new MockUsersService();
        m.bind(UsersService, toValue: usersService);
      });
    });

    afterEach(tearDownInjector);

    it('should redirect on success', async(inject(
      (UsersService us, UserRegisterController c, Window w) {
        var user = new User('name', 'email@example.com', 'password');
        c..displayName = user.displayName
         ..email = user.email
         ..password = user.password;

        when(us.registerUser(user)).thenReturn(new Future.value(new RegisterSuccess(10)));

        c.registerSubmit();
        microLeap();
      }
    )));

    it('should display registration error message', async(inject(
      (UsersService us, UserRegisterController c) {
        when(us.registerUser(any))
          .thenReturn(new Future.error(new RegistrationException('error', 'error message')));

        c.registerSubmit();
        microLeap();

        expect(c.errorMessage).toBe('error message');
      }
    )));

    it('should clear the error message on submit', async(inject(
      (UsersService us, UserRegisterController c) {
          when(us.registerUser(any)).thenReturn(new Future.error('error'));
          c.errorMessage = 'message';

          c.registerSubmit();

          expect(c.errorMessage).toBe('');
      }
    )));

    it('should display a generic error message on unknown error', async(inject(
      (UsersService us, UserRegisterController c) {
        when(us.registerUser(any)).thenReturn(new Future.error("error"));

        c.registerSubmit();
        microLeap();

        expect(c.errorMessage).toBeNotNull();
        expect(c.errorMessage).toBeTruthy();
      }
    )));
  });
}

