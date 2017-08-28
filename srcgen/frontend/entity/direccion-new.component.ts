import { Component, OnInit } from '@angular/core';
import { Direccion } from '../../models/direccion';
import { DireccionService } from '../../services/direccion.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-direccion-new',
  templateUrl: './direccion-new.component.html',
  styleUrls: ['./direccion-new.component.css']
})
export class DireccionNewComponent implements OnInit {

  direccion: Direccion;
  display = false;
  constructor(private direccionService: DireccionService, private location: Location) { }

  ngOnInit() {
    this.direccion = new Direccion();
  }

  guardar(direccion: Direccion): void {

    this.direccionService.create(direccion);
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