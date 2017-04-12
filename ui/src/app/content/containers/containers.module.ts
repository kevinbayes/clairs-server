/**
 * Created by kevinbayes on 19/12/16.
 */
import { ContainersComponent } from './containers.component';
import { ContainersRootComponent } from './containers.root.component';


import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';

@NgModule({
    imports: [

        RouterModule,
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
