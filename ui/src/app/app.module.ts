import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import {APP_BASE_HREF, CommonModule} from '@angular/common';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { MaterialModule } from './shared/material.module';

import { appRoutingProviders, AppRoutingModule } from './app.routes';

import { ContainersModule } from './content/containers/containers.module'
import { RegistriesModule } from './content/registries/registries.module';

import { AppComponent } from './app.component';

import { HeaderComponent } from './content/shared/header/header.component';
import { environment } from '../environments/environment';
import {RegistriesService} from "./services/registries.service";
import {ContainersService} from "./services/containers.service";


import { MdDataTableModule } from 'ng2-md-datatable';
import {DashboardModule} from "./content/dashboard/dashboard.module";
import { AppStoreModule } from '../store/app-store.module';
import {SummariesService} from "./services/summaries.service";
import { DetailComponent } from './registries/registry/detail/detail.component';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
  ],
  imports: [
    BrowserAnimationsModule,
    BrowserModule,
    CommonModule,
    FormsModule,
    HttpModule,
    MaterialModule,

    AppRoutingModule,

    DashboardModule,
    ContainersModule,
    RegistriesModule,

    MdDataTableModule,
    AppStoreModule,
  ],
  providers: [
    appRoutingProviders,
    {provide: APP_BASE_HREF, useValue: ((environment.production) ? '/' : '/ui/')},

    RegistriesService,
    ContainersService,
    SummariesService,
    ],
  bootstrap: [AppComponent]
})
export class AppModule { }
