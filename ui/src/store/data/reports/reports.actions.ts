import { Action } from '@ngrx/store';
import { type } from '../../utils';
import { PackageReport } from './package-report.model';

export const ActionTypes = {
  LOAD_REPORTS: type('[Reports] Load Reports by Id'),
  LOAD_REPORTS_COMPLETE: type('[Reports] Load Reports Complete'),
};

export class LoadReportsAction implements Action {
  public type = ActionTypes.LOAD_REPORTS;

  constructor(public payload: string) { }
}

export class LoadReportsCompleteAction implements Action {
  public type = ActionTypes.LOAD_REPORTS_COMPLETE;

  constructor(public payload: PackageReport[]) { }
}

export type Actions
  = LoadReportsAction
  | LoadReportsCompleteAction;
