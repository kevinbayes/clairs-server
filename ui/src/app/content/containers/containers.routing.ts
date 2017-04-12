/**
 * Created by kevinbayes on 19/12/16.
 */
import { ContainersComponent } from './containers.component';
import { ContainersRootComponent } from './containers.root.component';
import { Routes, RouterModule } from '@angular/router';
import { ModuleWithProviders }   from '@angular/core';


export const containersRoutes: Routes = [
    {
        path: 'containers',
        component: ContainersRootComponent,
        children: [
            {
                path: '',
                component: ContainersComponent,
            }
        ],
    },
];

export const containersRouting: ModuleWithProviders = RouterModule.forChild(containersRoutes);
