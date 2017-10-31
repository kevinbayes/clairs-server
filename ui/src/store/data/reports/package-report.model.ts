import { Vulnerability } from './vulnerability.model';

export interface PackageReport {
  packageName: string;
  packageVersion: string;
  vulnerabilitiesBefore: Vulnerability[];
  vulnerabilitiesAfter: Vulnerability[];
  upgradeImpact: number;
}
