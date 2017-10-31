/**
 * Created by kevinbayes on 19/12/16.
 */
import { ContainersComponent } from './containers.component';
import { ContainersRootComponent } from './containers.root.component';


import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { MaterialModule } from '../../shared/material.module';
import {BrowserModule} from "@angular/platform-browser";

import { ReportsComponent } from './reports/reports.component';
import { ContainerComponent } from './container/container.component';
import {NewContainerModalComponent} from "./new/new.container.modal.component";
import {ReactiveFormsModule} from "@angular/forms";

@NgModule({
    imports: [

        RouterModule,
        MaterialModule,
        BrowserModule,
        ReactiveFormsModule,
    ],
    exports: [

      ContainersRootComponent,
    ],
    declarations: [
      ContainersRootComponent,

      ContainersComponent,

      ReportsComponent,

      ContainerComponent,

      NewContainerModalComponent,
    ],
    providers: [

    ],
    entryComponents: [
      NewContainerModalComponent,
    ]
})

export class ContainersModule { }
