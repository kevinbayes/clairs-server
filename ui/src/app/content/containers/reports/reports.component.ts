import { Component, OnInit } from '@angular/core';
import { PackageReport } from '../../../../store/data/reports/package-report.model';
import { Store } from '@ngrx/store';
import * as fromRoot from 'store';
import { Subscription } from "rxjs/Subscription";
import { LoadReportsAction } from '../../../../store/data/reports/reports.actions';
import { ActivatedRoute } from "@angular/router";
import { Vulnerability } from '../../../../store/data/reports/vulnerability.model';

@Component({
  selector: 'app-reports',
  templateUrl: './reports.component.html',
  styleUrls: ['./reports.component.less']
})
export class ReportsComponent implements OnInit {

  public reports: PackageReport[];
  private subscriptions: Subscription[] = [];

  constructor(private store: Store<fromRoot.State>, private route: ActivatedRoute) { }

  public ngOnInit() {
    this.subscriptions.push(this.route
      .params
      .subscribe(params => {
        if (params['id']) {
          this.store.dispatch(new LoadReportsAction(params['id']))
        }
      })
    );

    this.subscriptions.push(
      this.store
        .select(fromRoot.getCurrentReport)
        .subscribe(result => this.reports = result)
    );
  }

  public ngOnDestroy() {
    this.subscriptions.forEach(s => s.unsubscribe());
  }

  public getVulnerabilitySummary(vulnerabilities: Vulnerability[]): string {
    let highTotal = 0;
    let others = 0;

    vulnerabilities
      .filter(v => v.level == Vulnerability.VulnerabilityLevel.High)
      .forEach(v => highTotal += v.count);
    vulnerabilities
      .filter(v => v.level != Vulnerability.VulnerabilityLevel.High)
      .forEach(v => others += v.count);

    return `${highTotal} High + ${others} additional`;
  }

}
