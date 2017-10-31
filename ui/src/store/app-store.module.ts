import { NgModule } from '@angular/core';
import { StoreModule } from '@ngrx/store';
import { EffectsModule } from '@ngrx/effects';
import { StoreDevtoolsModule } from '@ngrx/store-devtools';
import { reducer } from './index';
import { RouterStoreModule } from "@ngrx/router-store";
import { ReportsEffects } from './data/reports/reports.effects';
import { ReportsService } from './data/reports/reports.service';

@NgModule({
  imports: [
    StoreModule.provideStore(reducer),
    RouterStoreModule.connectRouter(),
    StoreDevtoolsModule.instrumentOnlyWithExtension(),
    EffectsModule.run(ReportsEffects),
  ],
  providers: [
    ReportsService,
  ]
})
export class AppStoreModule { }
