/**
 * Created by kevinbayes on 19/12/16.
 */
import { ContainersComponent } from './containers.component';
import { ContainersRootComponent } from './containers.root.component';


import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { MaterialModule } from '../../shared/material.module';
import {BrowserModule} from "@angular/platform-browser";

import { MdDataTableModule } from 'ng2-md-datatable';
import { ReportsComponent } from './reports/reports.component';

@NgModule({
    imports: [

        RouterModule,
        MaterialModule,
        BrowserModule,
        MdDataTableModule,
    ],
    exports: [
      ContainersRootComponent,

      ContainersComponent,
    ],
    declarations: [
      ContainersRootComponent,

      ContainersComponent,

      ReportsComponent,
    ],
    providers: [

    ],
    entryComponents: [
    ]
})

export class ContainersModule { }
