library register_form_spec;

import 'dart:async';
import 'dart:html';

import 'package:guinness/guinness.dart';
import 'package:mockito/mockito.dart';
import 'package:mock/mock.dart' as m;
import 'package:di/di.dart';
import 'package:angular/angular.dart';
import 'package:angular/mock/module.dart';

import 'package:user_management/controller/register_form.dart';
import 'package:user_management/service/users.dart';
import 'package:user_management/model/user.dart';

@proxy
class MockUsersService extends Mock implements UsersService {
  noSuchMethod(Invocation invocation) => super.noSuchMethod(invocation);
}

void main() {
  describe('RegisterFormController', () {
    MockLocation _loc;

    beforeEach(() {
      setUpInjector();
      module((Module m) {
        m.bind(RegisterFormController);
        var usersService = new MockUsersService();
        m.bind(UsersService, toValue: usersService);
      });
    });

    beforeEach(inject((Window win) {
      _loc = (win as MockWindow).location;
      _loc.when(m.callsTo('get search')).alwaysReturn('?continue=%2Fpath');
    }));

    afterEach(tearDownInjector);

    it('should redirect to received url on success', async(inject(
      (UsersService us, RegisterFormController c) {
        var user = new User('name', 'email@example.com', 'password');
        c..displayName = user.displayName
         ..email = user.email
         ..password = user.password;

        var response = new RegisterResponse.empty()..valid = true
                                                   ..redirectUri = '/redirect';
        when(us.registerUser(user, any)).thenReturn(new Future.value(response));

        c.registerSubmit();
        microLeap();
        _loc.getLogs(m.callsTo('assign', '/redirect')).verify(m.happenedExactly(1));
      }
    )));

    it('should succeed with no continue parameter', async(inject(
      (UsersService us, RegisterFormController c) {
        _loc.reset();
        _loc.when(m.callsTo('get search')).alwaysReturn('');
        when(us.registerUser(any, any)).thenReturn(new Future.error('error'));

        c.registerSubmit();
        verify(us.registerUser(any, null)).called(1);
      }
    )));

    it('should set field validation errors', async(inject(
      (UsersService us, RegisterFormController c) {
        var errors = [new InputError('email', 'error1'), new InputError('display_error', 'error2')];
        when(us.registerUser(any, any))
          .thenReturn(new Future.value(new RegisterResponse.empty()..valid = false
                                                                   ..errors = errors));

        c.registerSubmit();
        microLeap();

        expect(c.inputErrors.length).toBe(2);
        expect(c.inputErrors[errors[0].name]).toBe('error1');
        expect(c.inputErrors[errors[1].name]).toBe('error2');
      }
    )));

    it('should clear errors on submit', async(inject(
      (UsersService us, RegisterFormController c) {
          when(us.registerUser(any, any)).thenReturn(new Future.error('error'));
          c.inputErrors = {'email': 'error'};
          c.errorMessage = 'message';

          c.registerSubmit();

          expect(c.errorMessage).toBe('');
          expect(c.inputErrors.length).toBe(0);
      }
    )));

    it('should set a generic error message on unknown error', async(inject(
      (UsersService us, RegisterFormController c) {
        when(us.registerUser(any, any)).thenReturn(new Future.error("error"));

        c.registerSubmit();
        microLeap();

        expect(c.errorMessage.length).toBeGreaterThan(0);
      }
    )));
  });
}

