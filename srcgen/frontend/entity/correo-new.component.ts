import { Component, OnInit } from '@angular/core';
import { Correo } from '../../models/correo';
import { CorreoService } from '../../services/correo.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-correo-new',
  templateUrl: './correo-new.component.html',
  styleUrls: ['./correo-new.component.css']
})
export class CorreoNewComponent implements OnInit {

  correo: Correo;
  display = false;
  constructor(private correoService: CorreoService, private location: Location) { }

  ngOnInit() {
    this.correo = new Correo();
  }

  guardar(correo: Correo): void {

    this.correoService.create(correo);
    this.display = true;

  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}