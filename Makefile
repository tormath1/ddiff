build:
	@go build -o ddiff main.go

download:
	wget https://qa-reports.gentoo.org/output/genrdeps/rdeps.tar.xz
	@mkdir rdeps
	tar -xf rdeps.tar.xz -C ./rdeps
	@rm rdeps.tar.xz
