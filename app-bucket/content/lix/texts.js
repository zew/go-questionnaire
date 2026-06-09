const CATS = [
    {
        id:    'steuern',
        label: 'Steuern',
        lblsh: 'Steuern   ',
        color: '#9D8FBF',
        tooltip: 'Laufende Besteuerung des Unternehmens, Besteuerung des Unternehmens im Erbschaftsfall, Komplexität des Steuersystems',
        subs: [
            { id: 'biz', label: 'Laufende Besteuerung des Unternehmens', hint: 'Einkommen-, Körperschafts- und Gewerbesteuer', tooltip: 'Besteuerung nationaler und grenzüberschreitender Geschäftstätigkeit, etwa durch Einkommen- und Körperschaftssteuer sowie Grundsteuer.' },
            { id: 'erb', label: 'Besteuerung des Unternehmens im Erbschaftsfall', hint: 'Höhe und Vergünstigungen für Unternehmen, Administration', tooltip: 'Höhe der Erbschaftssteuer, Vergünstigungen für Unternehmen und administrative Regelungen bei der Unternehmensübertragung.' },
            { id: 'kpl', label: 'Komplexität des Steuersystems', hint: 'Aufwand zur Erfüllung der Steuerpflichten und Sozialabgaben', tooltip: 'Benötigter Arbeitsaufwand zur Erfüllung aller steuerlichen Pflichten und Sozialabgaben.' }
        ]
    },
    {
        id:    'arbeit',
        label: 'Arbeitskräfte (Kosten & Produktivität)',
        lblsh: 'Arbeitskosten & -produktivität',
        color: '#7EAFA3',
        tooltip: 'Arbeitskosten (Löhne und Lohnnebenkosten pro Stunde), Produktivität und Bildungsstand der Arbeitskräfte',
        subs: [
            { id: 'kos', label: 'Arbeitskosten', hint: 'Löhne und Lohnnebenkosten pro Stunde', tooltip: 'Alle aus Arbeitgebersicht anfallenden Arbeitskosten pro Stunde.' },
            { id: 'pro', label: 'Produktivität und Bildungsstand der Arbeitskräfte', hint: 'BIP je Arbeitsstunde, Bildungsausgaben, Bildungsniveau', tooltip: 'BIP je gearbeitete Arbeitsstunde, öffentliche und private Bildungsausgaben, Bildungsniveau der erwerbsfähigen Bevölkerung.' }
        ]
    },
    {
        id:    'fin',
        label: 'Finanzierungsbedingungen für Unternehmen und Zustand der öffentlichen Finanzen',
        lblsh: 'Finanzierungsbedingungen & öffentl. Finanzen',
        color: '#8FAD7E',
        tooltip: 'Verfügbarkeit von Unternehmenskrediten, Durchsetzbarkeit von Kreditforderungen, Zustand der öffentlichen Finanzen und private Verschuldungssituation',
        subs: [
            { id: 'krd', label: 'Verfügbarkeit von Unternehmenskrediten, Durchsetzbarkeit von Kreditforderungen', hint: 'Kreditmarkt, Gläubigerschutz, Kreditinformation', tooltip: 'Größe des Marktes für Unternehmenskredite, Risikoanfälligkeit, Zugangsmöglichkeiten, rechtliche Stellung von Gläubigern und Schuldnern, Umfang und Qualität von Kreditinformationen' },
            { id: 'vrs', label: 'Zustand der öffentlichen Finanzen und private Verschuldungssituation', hint: 'Verschuldung öffentlicher und privater Haushalte, Sovereign Ratings', tooltip: 'Verschuldung öffentlicher und privater Haushalte, Kreditwürdigkeit des Landes gemäß Bewertung der Rating-Agenturen.' }
        ]
    },
    {
        id:    'reg',
        label: 'Bürokratie und Regulierung',
        lblsh: 'Bürokratie & Regulierung  ',
        color: '#C4A265',
        tooltip: 'Vorschriften in den Bereichen Arbeitsverträge und Außenhandel (z.B. Kündigungsschutz; Tarifverträge; Zollvorschriften); Allgemeine laufende Berichtspflichten und Vorschriften bei Unternehmensgründung; Betriebliche Mitbestimmung',
        subs: [
            { id: 'ins', label: 'Vorschriften in den Bereichen Arbeitsverträge und Außenhandel', hint: 'z.B. Kündigungsschutz; Tarifverträge; Zollvorschriften', tooltip: 'Arbeitsrechtliche Vorgaben, welche die Vertragsfreiheit begrenzen, z.B. bei Kündigungen und Arbeitszeiten; Regularien bei grenzüberschreitendem Handel.' },
            { id: 'vor', label: 'Allgemeine laufende Berichtspflichten und Vorschriften', hint: 'Aufwand zur Einhaltung der Vorschriften und Anforderungen im laufenden Betrieb und bei der Unternehmensgründung', tooltip: 'Aufwand die regierungsseitigen Vorschriften und behördlichen Anforderungen im laufenden Betrieb zu erfüllen; Aufwand an Formalitäten, Zeit und Kosten bei Unternehmensgründung' },
            { id: 'mit', label: 'Betriebliche Mitbestimmung', hint: 'Aufwand Arbeitnehmerentscheidungen einzubeziehen', tooltip: 'Aufwand und Kosten für Bereitstellung von Ressourcen sowie Einschränkungen und Verlangsamung unternehmerischer Entscheidungen.' }
        ]
    },
    {
        id:    'inf',
        label: 'Qualität von Infrastruktur und politischen Institutionen',
        lblsh: 'Infrastruktur & Institutionen  ',
        color: '#6E9BBF',
        tooltip: 'Transportinfrastruktur und digitale Infrastruktur; Rechtlich-institutionelle Rahmenbedingungen',
        subs: [
            { id: 'tra', label: 'Transport- und digitale Infrastruktur', hint: 'Transportinfrastruktur (Straße, Schiene, Luft), Informations- und Kommunikationsinfrastruktur', tooltip: 'Straßen-, Schienen-, Flugverkehr; Cybersicherheit, Leistungsfähigkeit Breitbandnetze und Mobilfunknetze.' },
            { id: 'rec', label: 'Politische Institutionen', hint: 'Rechtlich-institutionelle Rahmenbedingungen', tooltip: 'Rechtssicherheit, Kriminalität und politische Stabilität, Korruptionskontrolle.' }
        ]
    },
    {
        id:    'ene',
        label: 'Energiesystem und Energiepolitik',
        lblsh: 'Energie',
        color: '#BF7E85',
        tooltip: 'Energiepreis, Energiesicherheit, Klimapolitische Ambitionen',
        subs: [
            { id: 'pre', label: 'Energiepreise', hint: 'Strom-, Gas- und Kraftstoffpreise', tooltip: 'Höhe der Strom-, Gas- und Kraftstoffpreise als direkte Kostenkomponente für Unternehmen.' },
            { id: 'sic', label: 'Energiesicherheit', hint: 'Akute und langfristige Zuverlässigkeit der Energieversorgung', tooltip: 'Zuverlässigkeit der Versorgung, Anzahl und Dauer der Stromunterbrechungen; Risikobewertung der Energieimportländer, Zusammensetzung der Energieimportländer am Gesamtimport.' },
            { id: 'kli', label: 'Klimapolitische Ambitionen', hint: 'Zielerreichungsgrad bei der Reduktion der Treibhausgasemissionen', tooltip: 'Erwartete Energie- und Bürokratiekosten angesichts des Zielerreichungsgrads der gesetzten Klimaziele; Anpassungslasten für Unternehmen, die sich aus den klimapolitischen Zielen des Landes ergeben' }
        ]
    }
];


console.info("texts.js loaded")