import { ActionReducer, combineReducers } from '@ngrx/store';
import { compose } from '@ngrx/core/compose';
import { storeFreeze } from 'ngrx-store-freeze';
import { environment } from '../environments/environment';
import { createSelector} from 'reselect';
import { RouterState, routerReducer } from '@ngrx/router-store';

import * as fromReports from './data/reports/reports.reducers';

export interface State {
  reports: fromReports.State;
  router: RouterState;
};

const reducers = {
  reports: fromReports.reducer,
  router: routerReducer,
};

const developmentReducer: ActionReducer<State> = compose(storeFreeze, combineReducers)(reducers);
const productionReducer: ActionReducer<State> = combineReducers(reducers);

export function reducer(state: any, action: any) {
  if (environment.production) {
    return productionReducer(state, action);
  } else {
    return developmentReducer(state, action);
  }
}

export const getReportsState = (state: State) => state.reports;

export const getAllReports = createSelector(getReportsState, fromReports.getAllReports);
export const getCurrentReport = createSelector(getReportsState, fromReports.getCurrentReport);
