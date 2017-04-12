import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { MaterialModule } from './shared/material.module';

import { AppComponent } from './app.component';
import { ContainersComponent } from './content/containers/containers.component';
import { HeaderComponent } from './content/shared/header/header.component';
import { SidemenuComponent } from './content/shared/sidemenu/sidemenu.component';

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
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
