/**
 * Created by kevinbayes on 19/12/16.
 */
import { DashboardComponent } from './dashboard.component';
import { DashboardRootComponent } from './dashboard.root.component';


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
      DashboardRootComponent,

      DashboardComponent,
    ],
    declarations: [
      DashboardRootComponent,

      DashboardComponent,
    ],
    providers: [

    ],
    entryComponents: [
    ]
})

export class DashboardModule { }
