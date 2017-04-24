/**
 * Created by kevinbayes on 19/12/16.
 */
import { RegistriesComponent } from './registries.component';
import { RegistriesRootComponent } from './registries.root.component';
import { Routes, RouterModule } from '@angular/router';
import { ModuleWithProviders }   from '@angular/core';
import {RegistryComponent} from "./registry/registry.component";


export const registriesRoutes: Routes = [
    {
        path: 'registries',
        component: RegistriesRootComponent,
        children: [
            {
                path: '',
                component: RegistriesComponent,
            },
            {
                path: ':id',
                component: RegistryComponent,
            }
        ],
    },
];

export const registriesRouting: ModuleWithProviders = RouterModule.forChild(registriesRoutes);
