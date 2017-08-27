import { Component, OnInit } from '@angular/core';
import { Votante } from '../../models/votante';
import { VotanteService } from '../../services/votante.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-votante-new',
  templateUrl: './votante-new.component.html',
  styleUrls: ['./votante-new.component.css']
})
export class VotanteNewComponent implements OnInit {

  votante: Votante;
  display = false;
  constructor(private votanteService: VotanteService, private location: Location) { }

  ngOnInit() {
    this.votante = new Votante();
  }

  guardar(votante: Votante): void {

    this.votanteService.create(votante);
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