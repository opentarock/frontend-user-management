exports.config = {

  specs: [
    'test_e2e/**/*_spec.dart',
  ],

  baseUrl: 'http://localhost:' + (process.env.HTTP_PORT || '8080'),

  seleniumServerJar: "node_modules/protractor/selenium/selenium-server-standalone-2.42.2.jar",

  chromeDriver: 'node_modules/protractor/selenium/chromedriver',

  chromeOnly: false,

  multiCapabilities: [
    { 'browserName': 'chrome' }
  ],

  jasmineNodeOpts: {
    isVerbose: true,
    showColors: true,
    includeStackTrace: true,
    defaultTimeoutInterval: 50000
  }
};
