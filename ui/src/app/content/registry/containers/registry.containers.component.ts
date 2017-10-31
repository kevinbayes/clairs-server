import {AfterViewInit, Component, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {ContainersService} from "../../../services/containers.service";
import {Observable} from "rxjs/Observable";
import 'rxjs/add/observable/from';
import {Subject} from "rxjs/Subject";
import {RegistryNewContainerModalComponent} from "./new/registry.new.container.modal.component";
import {MatDialog} from "@angular/material";
import {RegistriesService} from "../../../services/registries.service";
import {ActivatedRoute, Route} from "@angular/router";

@Component({
  selector: 'app-registry-containers',
  templateUrl: './registry.containers.component.html',
  styleUrls: ['./registry.containers.component.less']
})
export class RegistryContainersComponent implements OnInit, AfterViewInit, OnDestroy {

  private id: number;

  public pagination: any = {
    Pages: 0,
    Page: 0,
    Size: 10
  };
  public containers: any[];
  private unmount$: Subject<void> = new Subject<void>();

  constructor(private containersService: ContainersService,
              private registriesService: RegistriesService,
              private route: ActivatedRoute,
              public dialog: MatDialog) { }

  ngOnInit() {

    this.route.params.subscribe(params => {

      this.id = +params['id']; // (+) converts string 'id' to a number

      this.registriesService.listContainers(this.id).subscribe((res) => {
        this.pagination = res.Meta;
        this.containers = res.Entities;
      }, (err) => {

        console.error(err);
      });
    });
  }

  ngAfterViewInit() {
  }

  openDialog() {
    let dialogRef = this.dialog.open(RegistryNewContainerModalComponent);
    dialogRef.afterClosed().subscribe(result => {
    });
  }

  ngOnDestroy() {
    this.unmount$.next();
    this.unmount$.complete();
  }

}
