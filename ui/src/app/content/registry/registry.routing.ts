/**
 * Created by kevinbayes on 19/12/16.
 */
import { RegistryRootComponent } from './registry.root.component';
import { Routes, RouterModule } from '@angular/router';
import { ModuleWithProviders }   from '@angular/core';
import { RegistryComponent } from "./registry.component";


export const registryRoutes: Routes = [
    {
        path: 'registries',
        component: RegistryRootComponent,
        children: [
            {
                path: ':id',
                component: RegistryComponent,
            }
        ],
    },
];

export const registryRouting: ModuleWithProviders = RouterModule.forChild(registryRoutes);
