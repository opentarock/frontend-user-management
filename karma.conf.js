module.exports = function(config) {
  config.set({
    basePath: '.',
    frameworks: ['dart-unittest'],

    files: [
      'test/**/*_spec.dart',
      'test/init_guinness.dart',
      {pattern: '**/*.dart', watched: true, included: false, served: true},
      'packages/browser/dart.js'
    ],

    autoWatch: true,
    singleRun: false,
    captureTimeout: 20000,
    browserNoActivityTimeout: 300000,

    plugins: [
      'karma-dart',
      'karma-chrome-launcher',
      'karma-firefox-launcher'
    ],

    browsers: ['Firefox', 'Chrome']
  });
};