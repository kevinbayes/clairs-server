import {AfterViewInit, Component, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {ContainersService} from "../../../services/containers.service";
import {Subject} from "rxjs/Subject";
import {Observable} from "rxjs/Observable";
import {ActivatedRoute} from "@angular/router";

@Component({
  selector: 'app-container',
  templateUrl: './container.component.html',
  styleUrls: ['./container.component.less']
})
export class ContainerComponent implements OnInit, AfterViewInit, OnDestroy {

  private id: number;
  public pagination: any = {
    Pages: 0,
    Page: 0,
    Size: 10
  };
  public container: any = {Tags:[]};
  private unmount$: Subject<void> = new Subject<void>();

  constructor(private route: ActivatedRoute, private containersService: ContainersService) { }

  ngOnInit() {

    this.route.params.subscribe(params => {

      this.id = +params['id']; // (+) converts string 'id' to a number

      this.containersService.get(this.id).subscribe((res) => {

        console.log(res);
        this.container = res.Entity;
      }, (err) => {

        console.error(err);
      });
    });
  }

  ngAfterViewInit() {

  }

  ngOnDestroy() {
    this.unmount$.next();
    this.unmount$.complete();
  }

}
