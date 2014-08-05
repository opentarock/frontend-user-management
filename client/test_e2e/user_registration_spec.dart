library register_spec;

import 'package:protractor/protractor_api.dart';

void main() {
  describe('User registration app', () {
    beforeEach(() {
      browser.get('/register?continue=/register?success');
    });

    it('should redirect after registration', () {
      element(by.model('email')).sendKeys('user@example.com');
      element(by.css('input#signUp')).click();
      expect(browser.getCurrentUrl()).toBe('http://localhost:8080/register?success');
    });

    it('should make validation messages visible', () {
      element(by.model('displayName')).sendKeys('Name');
      element(by.model('email')).sendKeys('error@example.com');
      element(by.model('password')).sendKeys('password');
      element(by.css('input#signUp')).click();
      expect(element(by.css('span.input-error')).isDisplayed()).toBeTruthy();
    });

    it('should not display any validation messages on page load', () {
      expect(element(by.css('span.input-error')).isDisplayed()).toBe(false);
    });

  });
}