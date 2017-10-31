/**
 * Created by kevinbayes on 19/12/16.
 */
import { DashboardComponent } from './dashboard.component';
import { DashboardRootComponent } from './dashboard.root.component';


import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { MaterialModule } from '../../shared/material.module';
import {BrowserModule} from "@angular/platform-browser";
import {NgxChartsModule} from "@swimlane/ngx-charts";
import { NewRegistryModalComponent } from "./new/new.registry.modal.component";
import {ReactiveFormsModule} from "@angular/forms";

@NgModule({
    imports: [

      RouterModule,
      MaterialModule,
      BrowserModule,
      ReactiveFormsModule,
      NgxChartsModule,
    ],
    exports: [
      DashboardRootComponent,

      DashboardComponent,
    ],
    declarations: [
      DashboardRootComponent,

      DashboardComponent,

      NewRegistryModalComponent,
    ],
    providers: [

    ],
    entryComponents: [
      NewRegistryModalComponent,
    ]
})

export class DashboardModule { }
