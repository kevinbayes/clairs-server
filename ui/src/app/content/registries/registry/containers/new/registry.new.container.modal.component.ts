import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators }    from '@angular/forms';
import {MdDialogRef} from "@angular/material";
import {ContainersService} from "../../../../../services/containers.service";
import {RegistriesService} from "../../../../../services/registries.service";
import {Containers} from "../../../../../../store/data/containers/container.model";
import NewContainer = Containers.NewContainer;

@Component({
  selector: 'app-new-container-modal',
  templateUrl: './registry.new.container.modal.component.html',
  styleUrls: ['./registry.new.container.modal.component.less']
})
export class RegistryNewContainerModalComponent implements OnInit {

  newContainerForm: FormGroup;
  errorMessage: string;
  registries = [
    {id: '4', name: 'Docker Hub'}
  ];

  constructor(
    private dialogRef: MdDialogRef<RegistryNewContainerModalComponent>,
    private containersService: ContainersService,
    private registriesService: RegistriesService,
    private builder: FormBuilder) {

    this.createForm();
  }

  ngOnInit() {
  }

  createForm() {

    this.newContainerForm = this.builder.group({
      image: ['', Validators.required ],
      registry: ['', Validators.required ],
    });
  }

  onSubmit() {

    if(this.newContainerForm.status == "VALID") {

      const original = this.newContainerForm.value;

      const container: NewContainer = new NewContainer();
      container.image = `${original.image}`;
      container.registry = parseInt(original.registry);

      this.containersService
        .save(container).subscribe((success) => {

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
