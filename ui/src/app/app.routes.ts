import { RouterModule, Routes } from '@angular/router';
import { NgModule, ModuleWithProviders } from '@angular/core';

import { containersRoutes } from './content/containers/containers.routing';
import {registriesRoutes} from "./content/registries/registries.routing";
import {dashboardRoutes} from "./content/dashboard/dashboard.routing";

const appRoutes: Routes = [

  {
    path: '',
    children: [
      ...dashboardRoutes,
      ...containersRoutes,
      ...registriesRoutes,
      { path: '**', redirectTo: '/dashboard' },
    ]
  },
];

export const appRoutingProviders: any[] = [
];

export const routing: ModuleWithProviders = RouterModule.forRoot(appRoutes);

@NgModule({
  imports: [routing],
  exports: [RouterModule],
})
export class AppRoutingModule { }
