/**
 * Created by kevinbayes on 19/12/16.
 */
import { ContainersComponent } from './containers.component';
import { ContainersRootComponent } from './containers.root.component';


import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { MaterialModule } from '../../shared/material.module';

@NgModule({
    imports: [

        RouterModule,
        MaterialModule,
    ],
    exports: [
      ContainersRootComponent,

      ContainersComponent,
    ],
    declarations: [
      ContainersRootComponent,

      ContainersComponent,
    ],
    providers: [

    ],
    entryComponents: [
    ]
})

export class ContainersModule { }
