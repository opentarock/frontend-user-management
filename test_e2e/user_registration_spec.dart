library register_spec;

import 'package:protractor/protractor_api.dart';

void main() {
  describe('User registration app', () {
    beforeEach(() {
      browser.get('register.html');
    });

    it('should make error message visible', () {
      element(by.model('displayName')).sendKeys('Name');
      element(by.model('email')).sendKeys('error@example.com');
      element(by.model('password')).sendKeys('password');
      element(by.css('input#signUp')).click();
      expect(element(by.css('#error')).isDisplayed()).toBeTruthy();
    });
  });
}