import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { APP_BASE_HREF } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { MaterialModule } from './shared/material.module';

import { AppComponent } from './app.component';
import { ContainersComponent } from './content/containers/containers.component';
import { HeaderComponent } from './content/shared/header/header.component';
import { SidemenuComponent } from './content/shared/sidemenu/sidemenu.component';
import { environment } from '../environments/environment';

@NgModule({
  declarations: [
    AppComponent,
    ContainersComponent,
    HeaderComponent,
    SidemenuComponent
  ],
  imports: [
    BrowserAnimationsModule,
    BrowserModule,
    FormsModule,
    HttpModule,
    MaterialModule,
  ],
  providers: [{provide: APP_BASE_HREF, useValue: ((environment.production) ? '/' : '/ui/')}],
  bootstrap: [AppComponent]
})
export class AppModule { }
