/*global __dirname */
'use strict';
 
var gulp = require('gulp');
 
var $ = require('gulp-load-plugins')({ lazy: true });
var args = require('yargs').argv;
var swaggerTSGenerator = require('swagger-ts-generator');
var request = require('request');
var source = require('vinyl-source-stream');
 
var config = require('./gulp.config');
 
//--------------- gulp tasks ---------------
 
gulp.task('default', ['show-help']); // Set default gulp tasks
gulp.task('show-help', $.taskListing);
 
gulp.task('gen', ['gen-webapi']);
gulp.task('gen-webapi', ['gen-webapi-download-swagger'], genWebapi);
gulp.task('gen-webapi-download-swagger', genWebapiDownloadSwagger);
 
//--------------- generator tasks ---------------
 
function genWebapi(done) {
    swaggerTSGenerator.generateTSFiles(config.swagger.swaggerFile, config.swagger.swaggerTSGeneratorOptions);
    done();
}
function genWebapiDownloadSwagger(done) {
    process.env['NODE_TLS_REJECT_UNAUTHORIZED'] = '0'; // Ignore 'UNABLE_TO_VERIFY_LEAF_SIGNATURE' authorization error
 
    return request
        .get({
            url: config.swagger.url,
            headers: {
                'User-Agent': 'request',
                'content-type': 'application/json'
            }
        })
        .pipe(customPlumber('Error gen-webapi-autorest'))
        .pipe(source(config.files.swaggerJson))
        .pipe($.streamify($.jsbeautifier(/*{ mode: 'VERIFY_AND_WRITE' }*/)))
        .pipe(gulp.dest(config.folders.swaggerFolder));
}
 
function customPlumber(errTitle) {
    return $.plumber({
        errorHandler: $.notify.onError({
            // Customizing error title
            title: errTitle || 'Error running Gulp',
            message: 'Error: <%= error.message %>',
            sound: 'Glass'
        })
    });
}
 
function log(msg) {
    $.util.log($.util.colors.yellow(msg));
}
