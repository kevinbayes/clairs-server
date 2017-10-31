import {AfterViewInit, Component, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {ContainersService} from "../../services/containers.service";
import {Observable} from "rxjs/Observable";
import 'rxjs/add/observable/from';
import {Subject} from "rxjs/Subject";
import {NewContainerModalComponent} from "./new/new.container.modal.component";
import {MatDialog} from "@angular/material";

@Component({
  selector: 'app-containers',
  templateUrl: './containers.component.html',
  styleUrls: ['./containers.component.less']
})
export class ContainersComponent implements OnInit, AfterViewInit, OnDestroy {

  public pagination: any = {
    Pages: 0,
    Page: 0,
    Size: 10
  };
  public containers: any[];
  private unmount$: Subject<void> = new Subject<void>();

  constructor(private containersService: ContainersService,
              public dialog: MatDialog) { }

  ngOnInit() {

    this.containersService.all().subscribe((res) => {

      console.log(res);
      this.pagination = res.Meta;
      this.containers = res.Entities;
    }, (err) => {

      console.error(err);
    });
  }

  ngAfterViewInit() {

  }

  openDialog() {
    let dialogRef = this.dialog.open(NewContainerModalComponent);
    dialogRef.afterClosed().subscribe(result => {
    });
  }

  ngOnDestroy() {
    this.unmount$.next();
    this.unmount$.complete();
  }

}
