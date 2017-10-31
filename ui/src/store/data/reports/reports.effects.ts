import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Rx';
import { Effect, Actions, toPayload } from '@ngrx/effects';
import { ActionTypes, LoadReportsCompleteAction } from './reports.actions';
import { Action } from '@ngrx/store';
import { ReportsService } from './reports.service';

@Injectable()
export class ReportsEffects {
    @Effect()
    public list$: Observable<Action> = this.actions$
        .ofType(ActionTypes.LOAD_REPORTS)
        .map(toPayload)
        .switchMap((id) => {
          return this.service
            .get(id)
            .map((results) => new LoadReportsCompleteAction(results))
            .catch(err => err); // add global exception handler for api
        });

    constructor(private actions$: Actions, private service: ReportsService) { }
}
