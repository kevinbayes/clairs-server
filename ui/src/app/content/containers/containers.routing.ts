/**
 * Created by kevinbayes on 19/12/16.
 */
import { ContainersComponent } from './containers.component';
import { ContainersRootComponent } from './containers.root.component';
import { Routes, RouterModule } from '@angular/router';
import { ModuleWithProviders }   from '@angular/core';

import { ContainerComponent } from './container/container.component';
import { ReportsComponent } from './reports/reports.component';


export const containersRoutes: Routes = [
    {
        path: 'containers',
        component: ContainersRootComponent,
        children: [
            {
                path: '',
                component: ContainersComponent,
            },
            {
                path: ':id',
                component: ContainerComponent,
            },
            {
                path: ':id/report/:tag',
                component: ReportsComponent,
            }
        ],
    },
];

export const containersRouting: ModuleWithProviders = RouterModule.forChild(containersRoutes);
