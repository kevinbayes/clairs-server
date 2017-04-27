/**
 * Created by kevinbayes on 19/12/16.
 */
import { RegistryRootComponent } from './registry.root.component';
import { ReactiveFormsModule } from '@angular/forms';

import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { MaterialModule } from '../../shared/material.module';
import {BrowserModule} from "@angular/platform-browser";
import {MdDataTableModule} from "ng2-md-datatable";
import { RegistryComponent } from './registry.component';
import {SidemenuComponent} from "../shared/sidemenu/sidemenu.component";
import {RegistryContainersComponent} from "./containers/registry.containers.component";
import {RegistryNewContainerModalComponent} from "./containers/new/registry.new.container.modal.component";
import {RegistryDetailComponent} from "./detail/detail.component";
import { DashboardComponent } from './dashboard/dashboard.component';

@NgModule({
    imports: [

      RouterModule,
      MaterialModule,
      BrowserModule,
      ReactiveFormsModule,
      MdDataTableModule,
    ],
    exports: [
      RegistryRootComponent,
    ],
    declarations: [
      SidemenuComponent,

      RegistryRootComponent,

      RegistryComponent,

      RegistryComponent,
      RegistryDetailComponent,
      RegistryContainersComponent,

      RegistryNewContainerModalComponent,
      DashboardComponent,
    ],
    providers: [

    ],
    entryComponents: [

      RegistryNewContainerModalComponent,
    ]
})

export class RegistryModule { }
