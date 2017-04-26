import { Component, OnInit } from '@angular/core';

import { FormBuilder, FormGroup, Validators }    from '@angular/forms';
import {RegistriesService} from "../../../../services/registries.service";

@Component({
  selector: 'app-registry-detail',
  templateUrl: './detail.component.html',
  styleUrls: ['./detail.component.less']
})
export class RegistryDetailComponent implements OnInit {

  registryForm: FormGroup;
  errorMessage: string;

  constructor(private builder: FormBuilder,
    private registriesService: RegistriesService,
  ) {

    this.createForm();
  }

  ngOnInit() {
  }

  createForm() {

    this.registryForm = this.builder.group({
      id: ['', Validators.required ],
      name: ['', Validators.required ],
      description: '',
      uri: ['', Validators.required ],
      username: '',
      password: '',
    });
  }

  onSubmit() {

    if(this.registryForm.status == "VALID") {

      const registry = this.registryForm.value;
      this.registriesService
        .update(registry).subscribe((success) => {

      }, (err) => {

        this.errorMessage = err.json().Message;
      });
    } else {

      this.errorMessage = "Please check your form.";
    }
  }

}
