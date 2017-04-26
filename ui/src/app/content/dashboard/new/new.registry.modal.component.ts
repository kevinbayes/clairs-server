import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators }    from '@angular/forms';
import { RegistriesService } from "../../../services/registries.service";
import {MdDialogRef} from "@angular/material";

@Component({
  selector: 'app-new-registry-modal',
  templateUrl: './new.registry.modal.component.html',
  styleUrls: ['./new.registry.modal.component.less']
})
export class NewRegistryModalComponent implements OnInit {

  newRegistryForm: FormGroup;
  errorMessage: string;

  constructor(
    private dialogRef: MdDialogRef<NewRegistryModalComponent>,
    private registriesService: RegistriesService,
    private builder: FormBuilder) {

    this.createForm();
  }

  ngOnInit() {
  }

  createForm() {

    this.newRegistryForm = this.builder.group({
      name: ['', Validators.required ],
      description: '',
      uri: ['', Validators.required ],
      username: '',
      password: '',
    });
  }

  onSubmit() {

    if(this.newRegistryForm.status == "VALID") {

      const registry = this.newRegistryForm.value;
      this.registriesService
        .save(registry).subscribe((success) => {

        this.dialogRef.close("SUCCESS");
      }, (err) => {

        this.errorMessage = err.json().Message;
      });
    } else {

      this.errorMessage = "Please check your form.";
    }
  }

  close() {

    this.dialogRef.close();
  }
}
