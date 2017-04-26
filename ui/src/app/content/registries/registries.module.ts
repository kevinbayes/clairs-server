/**
 * Created by kevinbayes on 19/12/16.
 */
import { RegistriesComponent } from './registries.component';
import { RegistriesRootComponent } from './registries.root.component';
import { ReactiveFormsModule } from '@angular/forms';

import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { MaterialModule } from '../../shared/material.module';
import {BrowserModule} from "@angular/platform-browser";
import {MdDataTableModule} from "ng2-md-datatable";
import { RegistryComponent } from './registry/registry.component';
import { NewRegistryModalComponent } from './new/new.registry.modal.component';
import {SidemenuComponent} from "../shared/sidemenu/sidemenu.component";
import {RegistryContainersComponent} from "./registry/containers/registry.containers.component";
import {RegistryNewContainerModalComponent} from "./registry/containers/new/registry.new.container.modal.component";
import {RegistryDetailComponent} from "./registry/detail/detail.component";

@NgModule({
    imports: [

      RouterModule,
      MaterialModule,
      BrowserModule,
      ReactiveFormsModule,
      MdDataTableModule,
    ],
    exports: [
      RegistriesRootComponent,
    ],
    declarations: [
      SidemenuComponent,

      RegistriesRootComponent,

      RegistriesComponent,

      RegistryComponent,
      RegistryDetailComponent,
      RegistryContainersComponent,

      NewRegistryModalComponent,
      RegistryNewContainerModalComponent,
    ],
    providers: [

    ],
    entryComponents: [
      NewRegistryModalComponent,
      RegistryNewContainerModalComponent,
    ]
})

export class RegistriesModule { }
