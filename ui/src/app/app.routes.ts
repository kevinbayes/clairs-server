import { RouterModule, Routes } from '@angular/router';
import { NgModule, ModuleWithProviders } from '@angular/core';

import { containersRoutes } from './content/containers/containers.routing';

const appRoutes: Routes = [

  {
    path: '',
    children: [
      ...containersRoutes,

      { path: '**', redirectTo: '/containers' },
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
