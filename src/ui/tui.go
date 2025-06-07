package ui

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	ltable "github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/lipgloss/tree"
	"github.com/vdbulcke/cert-monitor/src/certmonitor"
)

var warnStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("11"))

// var okStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))
var errorStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("9"))

type TableEntry struct {
	Key   string
	Value string
}

func (e *TableEntry) AsSlice() []string {
	return []string{e.Key, e.Value}
}

type TLSScan struct {
	SNI        string
	TlsVersion string
	Chain      []*x509.Certificate
}

func (s *TLSScan) String() string {
	if s.SNI != "" {
		return fmt.Sprintf("{SNI: '%s', TLS: '%s'}", s.SNI, s.TlsVersion)
	}

	return fmt.Sprintf("{TLS: '%s'}", s.TlsVersion)
}

func (u *CertMonitorUI) PrintScan(url string, resutls []*TLSScan, skew int) {

	t := tree.Root("SCAN")
	root := tree.Root(url)
	for _, res := range resutls {
		tr := tree.Root(res.String())

		curr := tr
		for _, c := range res.Chain {
			temp := tree.Root(u.TableX509Cert(c, skew))

			curr.Child(temp)
			curr = temp

		}

		root.Child(
			tr,
		)

	}
	t.Child(root)

	fmt.Println(t)
}

func fingerprint(cert *x509.Certificate) string {
	return fmt.Sprintf("%X", sha256.Sum256(cert.Raw))

}

func (u *CertMonitorUI) TableX509Cert(cert *x509.Certificate, skew int, extra ...*TableEntry) string {

	defaultStyles := table.DefaultStyles()

	warn := !IsValid(cert, skew)
	expired := !IsValid(cert, 0)
	data := [][]string{}
	data = append(data, []string{"subject", cert.Subject.String()})
	data = append(data, []string{"issuer", cert.Issuer.String()})
	data = append(data, []string{"not_before", cert.NotBefore.String()})
	exp := cert.NotAfter.String()
	data = append(data, []string{"not_after", exp})
	data = append(data, []string{"sha256 fingerprint", fingerprint(cert)})
	data = append(data, []string{"serial_number", cert.SerialNumber.String()})
	data = append(data, []string{"AKID", fmt.Sprintf("%X", cert.AuthorityKeyId)})
	data = append(data, []string{"SUKID", fmt.Sprintf("%X", cert.SubjectKeyId)})
	if len(cert.IssuingCertificateURL) > 0 {
		data = append(data, []string{"ca issuers", strings.Join(cert.IssuingCertificateURL, ",")})
	}
	if len(cert.CRLDistributionPoints) > 0 {
		data = append(data, []string{"crls", strings.Join(cert.CRLDistributionPoints, ",")})
	}
	if len(cert.OCSPServer) > 0 {
		data = append(data, []string{"ocsp", strings.Join(cert.OCSPServer, ",")})
	}

	if len(cert.DNSNames) > 0 {
		data = append(data, []string{"SAN dns", strings.Join(cert.DNSNames, ",")})
	}

	for _, e := range extra {
		data = append(data, e.AsSlice())
	}

	table := ltable.New().
		// Headers(columnNames...).
		Rows(data...).
		Width(150).
		StyleFunc(func(row, col int) lipgloss.Style {
			return defaultStyles.Cell
		})

	switch {
	case expired:
		return errorStyle.Render(table.Render())
	case warn:
		return warnStyle.Render(table.Render())

	default:
		return table.Render()

	}

}
func (u *CertMonitorUI) TablePrintX509Cert(cert *x509.Certificate, skew int, extra ...*TableEntry) {

	fmt.Println(u.TableX509Cert(cert, skew, extra...))
	// print PEM format to stdout
	err := pem.Encode(os.Stdout, &pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw})
	if err != nil {
		u.Logger.Error("Error writing PEM certifcate", "err", err)
		// return err
	}
}

func (u *CertMonitorUI) TablePrintJwkCert(jwk *certmonitor.CertMonitorJWK, skew int) {

	extra := []*TableEntry{

		{Key: "kid", Value: jwk.Kid},
		{Key: "alg", Value: jwk.Alg},
		{Key: "kty", Value: jwk.Kty},
	}

	for _, c := range jwk.Certs {
		u.TablePrintX509Cert(c, skew, extra...)
	}

}
