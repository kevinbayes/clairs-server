import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Rx';
import { PackageReport } from './package-report.model';
import { Vulnerability } from './vulnerability.model';
@Injectable()
export class ReportsService {
  constructor() { }

  public get(id: string): Observable<PackageReport[]> {
    return Observable.of(this.getMockReports());
  }

  private getMockReports(): PackageReport[] {
    return [
      {
        packageName: 'package1',
        packageVersion: 'latest',
        vulnerabilitiesBefore: [
          {
            level: Vulnerability.VulnerabilityLevel.High,
            count: 3,
          },
          {
            level: Vulnerability.VulnerabilityLevel.Low,
            count: 2,
          },
          {
            level: Vulnerability.VulnerabilityLevel.Negligible,
            count: 3,
          }
        ],
        vulnerabilitiesAfter: [
          {
            level: Vulnerability.VulnerabilityLevel.High,
            count: 1,
          }
        ],
        upgradeImpact: 2,
      },
      {
        packageName: 'package2',
        packageVersion: 'latest',
        vulnerabilitiesBefore: [],
        vulnerabilitiesAfter: [],
        upgradeImpact: 0,
      },
      {
        packageName: 'package3',
        packageVersion: 'latest',
        vulnerabilitiesBefore: [
          {
            level: Vulnerability.VulnerabilityLevel.High,
            count: 5,
          },
          {
            level: Vulnerability.VulnerabilityLevel.Low,
            count: 1,
          },
          {
            level: Vulnerability.VulnerabilityLevel.Negligible,
            count: 2,
          }
        ],
        vulnerabilitiesAfter: [
          {
            level: Vulnerability.VulnerabilityLevel.Negligible,
            count: 1,
          }
        ],
        upgradeImpact: 1,
      },
    ]
  }
}
