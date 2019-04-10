GODOC=godoc -index -links=true -notes="BUG|TODO|XXX|ISSUE"

docs:
	@echo "\nGenerating HTML files ...\n"
	@mkdir -p $(DOC_DIR)/cmd/zylisp $(DOC_DIR)/core $(DOC_DIR)/generator \
					 $(DOC_DIR)/generator/helpers $(DOC_DIR)/lexer \
					 $(DOC_DIR)/parser $(DOC_DIR)/repl $(DOC_DIR)/common \
					 $(DOC_DIR)/util
	@$(GODOC) -url /pkg/$(PROJ) > \
		$(DOC_DIR)/index.html
	@$(GODOC) -url /pkg/$(PROJ)/cmd/ > \
		$(DOC_DIR)/cmd/index.html
	@$(GODOC) -url /pkg/$(PROJ)/cmd/zylisp > \
		$(DOC_DIR)/cmd/zylisp/index.html
	@$(GODOC) -url /pkg/$(PROJ)/core > \
		$(DOC_DIR)/core/index.html
	@$(GODOC) -url /pkg/$(PROJ)/generator > \
		$(DOC_DIR)/generator/index.html
	@$(GODOC) -url /pkg/$(PROJ)/generator/helpers > \
		$(DOC_DIR)/generator/helpers/index.html
	@$(GODOC) -url /pkg/$(PROJ)/lexer > \
		$(DOC_DIR)/lexer/index.html
	@$(GODOC) -url /pkg/$(PROJ)/parser > \
		$(DOC_DIR)/parser/index.html
	@$(GODOC) -url /pkg/$(PROJ)/repl > \
		$(DOC_DIR)/repl/index.html
	@$(GODOC) -url /pkg/$(PROJ)/common > \
		$(DOC_DIR)/common/index.html

view-docs: docs
	@echo "\nView project docs in a browser at:"
	@echo "  http://localhost:6060/pkg/"
	@echo "In particular, the zylisp command docs are here:"
	@echo "  http://localhost:6060/pkg/$(PROJ)/cmd/zylisp/\n"
	@echo "Starting docs HTTP server ..."
	@GOPATH=`pwd` $(GODOC) -http=:6060 -goroot=`pwd`/doc

publish-docs: docs
	@cd doc && \
	git commit -am "Regen'ed docs." && \
	git push origin gh-pages
