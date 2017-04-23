/**
 * Created by kevinbayes on 19/12/16.
 */
import { RegistriesComponent } from './registries.component';
import { RegistriesRootComponent } from './registries.root.component';


import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { MaterialModule } from '../../shared/material.module';
import {BrowserModule} from "@angular/platform-browser";
import {MdDataTableModule} from "ng2-md-datatable";

@NgModule({
    imports: [

      RouterModule,
      MaterialModule,
      BrowserModule,
      MdDataTableModule,
    ],
    exports: [
      RegistriesRootComponent,

      RegistriesComponent,
    ],
    declarations: [
      RegistriesRootComponent,

      RegistriesComponent,
    ],
    providers: [

    ],
    entryComponents: [
    ]
})

export class RegistriesModule { }
