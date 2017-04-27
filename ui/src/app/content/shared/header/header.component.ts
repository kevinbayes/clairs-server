import { Component, OnInit } from '@angular/core';
import { MdIconRegistry } from '@angular/material';
import {DomSanitizer} from "@angular/platform-browser";

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.less']
})
export class HeaderComponent implements OnInit {

  private selectedRegistry: any;

  constructor(iconRegistry: MdIconRegistry, sanitizer: DomSanitizer) {
    iconRegistry.addSvgIcon(
      'github-circle',
        sanitizer.bypassSecurityTrustResourceUrl('assets/icons/github-circle.svg'));
  }

  ngOnInit() {
  }

}
