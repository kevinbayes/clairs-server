webpackJsonp([1,4],{

/***/ 150:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(15)();
// imports


// module
exports.push([module.i, "md-sidenav {\n  width: 200px;\n}\nhtml,\nbody,\nmaterial-app,\nmd-sidenav-container,\n.clair__content {\n  margin: 0;\n  width: 100%;\n  height: 100%;\n}\n", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ 151:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(15)();
// imports


// module
exports.push([module.i, "", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ 152:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(15)();
// imports


// module
exports.push([module.i, ".clair__fill-remaining-space {\n  -webkit-box-flex: 1;\n      -ms-flex: 1 1 auto;\n          flex: 1 1 auto;\n}\n.mat-toolbar {\n  box-shadow: 0 3px 4px 0 rgba(0, 0, 0, 0.14), 0 3px 3px -2px rgba(0, 0, 0, 0.12), 0 1px 8px 0 rgba(0, 0, 0, 0.2);\n  margin-bottom: 8px;\n}\n", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ 153:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(15)();
// imports


// module
exports.push([module.i, "", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ 155:
/***/ (function(module, exports) {

module.exports = "<app-header></app-header>\n<md-sidenav-container>\n  <md-sidenav #sidedrawer mode=\"side\" opened=\"true\">\n    <app-sidemenu></app-sidemenu>\n  </md-sidenav>\n  <div class=\"clair__content\">\n\n      Content\n  </div>\n</md-sidenav-container>\n"

/***/ }),

/***/ 156:
/***/ (function(module, exports) {

module.exports = "Working\n"

/***/ }),

/***/ 157:
/***/ (function(module, exports) {

module.exports = "<md-toolbar color=\"primary\">\n  <span>Clair</span>\n\n  <span class=\"clair__fill-remaining-space\"></span>\n\n  <span>\n    <a md-icon-button href=\"https://github.com/kevinbayes/clairs-server\" target=\"_blank\"><md-icon svgIcon=\"github-circle\"></md-icon></a>\n  </span>\n</md-toolbar>\n"

/***/ }),

/***/ 158:
/***/ (function(module, exports) {

module.exports = "<p>\n  sidemenu works!\n</p>\n"

/***/ }),

/***/ 210:
/***/ (function(module, exports, __webpack_require__) {

module.exports = __webpack_require__(83);


/***/ }),

/***/ 59:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return environment; });
var environment = {
    production: true
};
//# sourceMappingURL=environment.js.map

/***/ }),

/***/ 82:
/***/ (function(module, exports) {

function webpackEmptyContext(req) {
	throw new Error("Cannot find module '" + req + "'.");
}
webpackEmptyContext.keys = function() { return []; };
webpackEmptyContext.resolve = webpackEmptyContext;
module.exports = webpackEmptyContext;
webpackEmptyContext.id = 82;


/***/ }),

/***/ 83:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
Object.defineProperty(__webpack_exports__, "__esModule", { value: true });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__(3);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__angular_platform_browser_dynamic__ = __webpack_require__(88);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__app_app_module__ = __webpack_require__(91);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_3__environments_environment__ = __webpack_require__(59);




if (__WEBPACK_IMPORTED_MODULE_3__environments_environment__["a" /* environment */].production) {
    __webpack_require__.i(__WEBPACK_IMPORTED_MODULE_0__angular_core__["a" /* enableProdMode */])();
}
__webpack_require__.i(__WEBPACK_IMPORTED_MODULE_1__angular_platform_browser_dynamic__["a" /* platformBrowserDynamic */])().bootstrapModule(__WEBPACK_IMPORTED_MODULE_2__app_app_module__["a" /* AppModule */]);
//# sourceMappingURL=main.js.map

/***/ }),

/***/ 90:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__(3);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return AppComponent; });
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};

var AppComponent = (function () {
    function AppComponent() {
        this.title = 'app works!';
    }
    return AppComponent;
}());
AppComponent = __decorate([
    __webpack_require__.i(__WEBPACK_IMPORTED_MODULE_0__angular_core__["I" /* Component */])({
        selector: 'app-root',
        template: __webpack_require__(155),
        styles: [__webpack_require__(150)]
    })
], AppComponent);

//# sourceMappingURL=app.component.js.map

/***/ }),

/***/ 91:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_platform_browser__ = __webpack_require__(12);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__angular_core__ = __webpack_require__(3);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__angular_common__ = __webpack_require__(23);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_3__angular_forms__ = __webpack_require__(56);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_4__angular_http__ = __webpack_require__(57);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_5__angular_platform_browser_animations__ = __webpack_require__(89);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_6__shared_material_module__ = __webpack_require__(95);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_7__app_component__ = __webpack_require__(90);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_8__content_containers_containers_component__ = __webpack_require__(92);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_9__content_shared_header_header_component__ = __webpack_require__(93);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_10__content_shared_sidemenu_sidemenu_component__ = __webpack_require__(94);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_11__environments_environment__ = __webpack_require__(59);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return AppModule; });
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};












var AppModule = (function () {
    function AppModule() {
    }
    return AppModule;
}());
AppModule = __decorate([
    __webpack_require__.i(__WEBPACK_IMPORTED_MODULE_1__angular_core__["b" /* NgModule */])({
        declarations: [
            __WEBPACK_IMPORTED_MODULE_7__app_component__["a" /* AppComponent */],
            __WEBPACK_IMPORTED_MODULE_8__content_containers_containers_component__["a" /* ContainersComponent */],
            __WEBPACK_IMPORTED_MODULE_9__content_shared_header_header_component__["a" /* HeaderComponent */],
            __WEBPACK_IMPORTED_MODULE_10__content_shared_sidemenu_sidemenu_component__["a" /* SidemenuComponent */]
        ],
        imports: [
            __WEBPACK_IMPORTED_MODULE_5__angular_platform_browser_animations__["a" /* BrowserAnimationsModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_platform_browser__["a" /* BrowserModule */],
            __WEBPACK_IMPORTED_MODULE_3__angular_forms__["a" /* FormsModule */],
            __WEBPACK_IMPORTED_MODULE_4__angular_http__["a" /* HttpModule */],
            __WEBPACK_IMPORTED_MODULE_6__shared_material_module__["a" /* MaterialModule */],
        ],
        providers: [{ provide: __WEBPACK_IMPORTED_MODULE_2__angular_common__["a" /* APP_BASE_HREF */], useValue: ((__WEBPACK_IMPORTED_MODULE_11__environments_environment__["a" /* environment */].production) ? '/' : '/ui/') }],
        bootstrap: [__WEBPACK_IMPORTED_MODULE_7__app_component__["a" /* AppComponent */]]
    })
], AppModule);

//# sourceMappingURL=app.module.js.map

/***/ }),

/***/ 92:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__(3);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return ContainersComponent; });
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};

var ContainersComponent = (function () {
    function ContainersComponent() {
    }
    ContainersComponent.prototype.ngOnInit = function () {
    };
    return ContainersComponent;
}());
ContainersComponent = __decorate([
    __webpack_require__.i(__WEBPACK_IMPORTED_MODULE_0__angular_core__["I" /* Component */])({
        selector: 'app-containers',
        template: __webpack_require__(156),
        styles: [__webpack_require__(151)]
    }),
    __metadata("design:paramtypes", [])
], ContainersComponent);

//# sourceMappingURL=containers.component.js.map

/***/ }),

/***/ 93:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__(3);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__angular_material__ = __webpack_require__(58);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__angular_platform_browser__ = __webpack_require__(12);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return HeaderComponent; });
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};



var HeaderComponent = (function () {
    function HeaderComponent(iconRegistry, sanitizer) {
        iconRegistry.addSvgIcon('github-circle', sanitizer.bypassSecurityTrustResourceUrl('assets/icons/github-circle.svg'));
    }
    HeaderComponent.prototype.ngOnInit = function () {
    };
    return HeaderComponent;
}());
HeaderComponent = __decorate([
    __webpack_require__.i(__WEBPACK_IMPORTED_MODULE_0__angular_core__["I" /* Component */])({
        selector: 'app-header',
        template: __webpack_require__(157),
        styles: [__webpack_require__(152)]
    }),
    __metadata("design:paramtypes", [typeof (_a = typeof __WEBPACK_IMPORTED_MODULE_1__angular_material__["o" /* MdIconRegistry */] !== "undefined" && __WEBPACK_IMPORTED_MODULE_1__angular_material__["o" /* MdIconRegistry */]) === "function" && _a || Object, typeof (_b = typeof __WEBPACK_IMPORTED_MODULE_2__angular_platform_browser__["e" /* DomSanitizer */] !== "undefined" && __WEBPACK_IMPORTED_MODULE_2__angular_platform_browser__["e" /* DomSanitizer */]) === "function" && _b || Object])
], HeaderComponent);

var _a, _b;
//# sourceMappingURL=header.component.js.map

/***/ }),

/***/ 94:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__(3);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return SidemenuComponent; });
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};

var SidemenuComponent = (function () {
    function SidemenuComponent() {
    }
    SidemenuComponent.prototype.ngOnInit = function () {
    };
    return SidemenuComponent;
}());
SidemenuComponent = __decorate([
    __webpack_require__.i(__WEBPACK_IMPORTED_MODULE_0__angular_core__["I" /* Component */])({
        selector: 'app-sidemenu',
        template: __webpack_require__(158),
        styles: [__webpack_require__(153)]
    }),
    __metadata("design:paramtypes", [])
], SidemenuComponent);

//# sourceMappingURL=sidemenu.component.js.map

/***/ }),

/***/ 95:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_material__ = __webpack_require__(58);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__angular_core__ = __webpack_require__(3);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return MaterialModule; });
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
/**
 * Created by kevinbayes on 13/04/17.
 */


var MaterialModule = (function () {
    function MaterialModule() {
    }
    return MaterialModule;
}());
MaterialModule = __decorate([
    __webpack_require__.i(__WEBPACK_IMPORTED_MODULE_1__angular_core__["b" /* NgModule */])({
        imports: [
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["a" /* MdButtonModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["b" /* MdCheckboxModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["c" /* MdToolbarModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["d" /* MdCardModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["e" /* MdMenuModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["f" /* MdIconModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["g" /* MdInputModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["h" /* MdGridListModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["i" /* MdListModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["j" /* MdProgressSpinnerModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["k" /* MdSnackBarModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["l" /* MdRippleModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["m" /* MdSidenavModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["n" /* MdTabsModule */]
        ],
        exports: [
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["a" /* MdButtonModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["b" /* MdCheckboxModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["c" /* MdToolbarModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["d" /* MdCardModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["e" /* MdMenuModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["f" /* MdIconModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["g" /* MdInputModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["h" /* MdGridListModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["i" /* MdListModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["j" /* MdProgressSpinnerModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["k" /* MdSnackBarModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["l" /* MdRippleModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["m" /* MdSidenavModule */],
            __WEBPACK_IMPORTED_MODULE_0__angular_material__["n" /* MdTabsModule */]
        ]
    })
], MaterialModule);

//# sourceMappingURL=material.module.js.map

/***/ })

},[210]);
//# sourceMappingURL=main.bundle.js.map