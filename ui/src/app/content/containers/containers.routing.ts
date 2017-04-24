/**
 * Created by kevinbayes on 19/12/16.
 */
import { ContainersComponent } from './containers.component';
import { ContainersRootComponent } from './containers.root.component';
import { Routes, RouterModule } from '@angular/router';
import { ModuleWithProviders }   from '@angular/core';
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
                path: ':id/report',
                component: ReportsComponent,
            }
        ],
    },
];

export const containersRouting: ModuleWithProviders = RouterModule.forChild(containersRoutes);
