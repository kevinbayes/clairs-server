import { Actions, ActionTypes, LoadReportsCompleteAction, LoadReportsAction } from './reports.actions';
import { PackageReport } from './package-report.model';

export interface State {
  currentReportId: string | undefined;
  reports: { [key: string]: PackageReport[] };
}

export const initialState: State = {
  currentReportId: undefined,
  reports: {},
};

export function reducer(state = initialState, action: Actions): State {
  switch (action.type) {
    case ActionTypes.LOAD_REPORTS:
      return {
        currentReportId: (action as LoadReportsAction).payload,
        reports: state.reports,
      }
    case ActionTypes.LOAD_REPORTS_COMPLETE:
      let result = (action as LoadReportsCompleteAction).payload;

      return {
        currentReportId: state.currentReportId,
        reports:  Object.assign({}, state.reports, {
          [state.currentReportId]: result
        })
      };
    default:
      return state;
  }
}

export const getAllReports = (state: State) => state.reports;
export const getCurrentReport = (state: State) => state.reports[state.currentReportId];
