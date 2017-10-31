/**
 * Created by kevinbayes on 19/12/16.
 */
import { DashboardComponent } from './dashboard.component';
import { DashboardRootComponent } from './dashboard.root.component';
import { Routes, RouterModule } from '@angular/router';
import { ModuleWithProviders }   from '@angular/core';


export const dashboardRoutes: Routes = [
    {
        path: 'dashboard',
        component: DashboardRootComponent,
        children: [
            {
                path: '',
                component: DashboardComponent,
            }
        ],
    },
];

export const dashboardRouting: ModuleWithProviders = RouterModule.forChild(dashboardRoutes);
