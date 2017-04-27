import { Component, OnInit } from '@angular/core';
import {SummariesService} from "../../services/summaries.service";
import {RegistriesService} from "../../services/registries.service";
import {MdDialog} from "@angular/material";
import {NewRegistryModalComponent} from "./new/new.registry.modal.component";

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.less']
})
export class DashboardComponent implements OnInit {

  public pagination: any = {
    Pages: 0,
    Page: 0,
    Size: 11
  };
  public registries: any[] = [];

  view: any[] = [700, 150];

  scheme = "air";

  numbers = [
    {
      "name": "Registries",
      "value": 0
    },
    {
      "name": "Containers",
      "value": 0
    },
    {
      "name": "Reports",
      "value": 0
    }
  ];

  constructor(private summariesService: SummariesService,
              private registriesService: RegistriesService,
              public dialog: MdDialog) { }

  ngOnInit() {

    this.loadRegistries(0, 11);

    this.summariesService.overall()
      .subscribe((success) => {

        this.numbers = [
          {
            "name": "Registries",
            "value": success.Entity.Registries.Total
          },
          {
            "name": "Containers",
            "value": success.Entity.Containers.Total
          },
          {
            "name": "Reports",
            "value": success.Entity.Reports.Total
          }
        ];
      }, (err) => {

      });
  }

  openNewRegistry() {
    let dialogRef = this.dialog.open(NewRegistryModalComponent);
    dialogRef.afterClosed().subscribe(result => {
    });
  }

  loadRegistries(page, size) {
    this.registriesService.all(page, size).subscribe((res) => {

      console.log(res);
      this.pagination = res.Meta;
      if(res.Entities) {
        this.registries = this.registries.concat(res.Entities);
      }
    }, (err) => {

      console.error(err);
    });
  }

  loadNextPageofRegistries() {
    this.loadRegistries(this.pagination.Page + 1, this.pagination.Size);
  }

}
